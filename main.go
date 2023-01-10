package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var apikey string

func post(c *gin.Context) {
	headers, err := json.Marshal(c.Request.Header)
	if err != nil {
		return
	}

	log.Info().
		Str("client_ip", c.ClientIP()).
		Str("request_method", c.Request.Method).
		Str("path", c.Request.RequestURI).
		Int("status", c.Writer.Status()).
		Str("http_referrer", c.Request.Referer()).
		RawJSON("headers", headers).
		Msg("request")

	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		log.Error().Err(err).Msg("body_read_fail")
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	log.Info().Bytes("request_body", body).Msg("request_body")

	client := &http.Client{
		Timeout: 10 * time.Second,
	}

	req, err := http.NewRequest("POST", "https://api.openai.com/v1/completions", bytes.NewBuffer(body))
	if err != nil {
		log.Error().Err(err).Msg("upstream_request_prepare_fail")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", apikey)

	start := time.Now()

	res, err := client.Do(req)
	if err != nil {
		log.Error().Err(err).Msg("upstream_request_fail")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	defer res.Body.Close()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Error().Err(err).Msg("upstream_body_read_fail")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	end := time.Now()
	duration := end.Sub(start).Seconds()

	resString := string(resBody)

	if res.StatusCode != http.StatusOK {
		log.Error().
			Str("error", fmt.Sprintf("Upstream request failed due to status code %d", res.StatusCode)).
			Str("response_body", resString).
			Int("status", res.StatusCode).
			Float64("upstream_response_time", duration).
			Msg("upstream_bad_status")
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	log.Info().
		Str("response_body", resString).
		Int("status", res.StatusCode).
		Float64("upstream_response_time", duration).
		Msg("upstream_response")

	c.String(http.StatusOK, resString)
}

func main() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stdout)
	log.Logger = logger

	key, ok := os.LookupEnv("OpenAI_apikey")
	if !ok {
		log.Fatal().Msg("OpenAI_apikey ENV variable not found")
		os.Exit(1)
	}
	apikey = fmt.Sprintf("Bearer %s", key)

	router := gin.New()
	router.Use(gin.Recovery())
	router.SetTrustedProxies(nil)

	router.POST("/", post)

	const url = "localhost"
	const port = 8080
	log.Info().Int16("port", port).Str("url", url).Msg("start")
	router.Run(fmt.Sprintf("%s:%d", url, port))
}

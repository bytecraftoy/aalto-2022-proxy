package fi.bytecraft.proxy

import org.scalatra._
import scala.io.Source

class APIKeyProxyServlet extends ScalatraServlet {

  post("/") {
    val promptJson: String = request.body
    val apiKey: String = Source.fromFile("apikey.txt").getLines.mkString
    val timeout: Int = 10000

    val r = requests.post(
       "https://api.openai.com/v1/completions",
       data = promptJson,
       headers = Map("Authorization" -> apiKey, "Content-Type" -> "application/json"),
       readTimeout = timeout
    )

    r.text
  }
}

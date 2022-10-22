package fi.bytecraft.proxy

import org.scalatra._
import scala.io.Source

class APIKeyProxyServlet extends ScalatraServlet {

  post("/") {
    val promptJson: String = request.body
    val apiKey: String = sys.env.get("OpenAI_apikey").getOrElse("Apikey not received from environment variable.")
    val timeout: Int = 10000

    println("Apikey: " + sys.env)

    val r = requests.post(
       "https://api.openai.com/v1/completions",
       data = promptJson,
       headers = Map("Authorization" -> apiKey, "Content-Type" -> "application/json"),
       readTimeout = timeout
    )

    r.text
  }
}

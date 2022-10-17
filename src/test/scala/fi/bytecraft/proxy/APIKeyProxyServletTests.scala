package fi.bytecraft.proxy

import org.scalatra.test.scalatest._

class APIKeyProxyServletTests extends ScalatraFunSuite {

  addServlet(classOf[APIKeyProxyServlet], "/*")

}

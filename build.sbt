val ScalatraVersion = "2.8.2"

ThisBuild / scalaVersion := "2.13.9"
ThisBuild / organization := "fi.bytecraft"

lazy val hello = (project in file("."))
  .settings(
    name := "APIKeyProxy",
    version := "0.1.0-SNAPSHOT",
    assembly / assemblyOutputPath := new sbt.File("apikeyproxy.jar"),
    libraryDependencies ++= Seq(
      "org.scalatra" %% "scalatra" % ScalatraVersion,
      "org.scalatra" %% "scalatra-scalatest" % ScalatraVersion % "test",
      "ch.qos.logback" % "logback-classic" % "1.2.3" % "runtime",
      "org.eclipse.jetty" % "jetty-webapp" % "9.4.43.v20210629" % "container;compile",
      "javax.servlet" % "javax.servlet-api" % "3.1.0" % "provided",
      "com.lihaoyi" %% "requests" % "0.7.1",
      "com.lihaoyi" %% "upickle" % "1.6.0"
    )
  )

enablePlugins(SbtTwirl)
enablePlugins(JettyPlugin)

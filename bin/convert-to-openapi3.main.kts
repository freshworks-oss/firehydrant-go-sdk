#!/usr/bin/env kotlin

@file:DependsOn("io.swagger.parser.v3:swagger-parser:2.1.22")
@file:DependsOn("io.swagger:swagger-inflector:2.0.12")

import io.swagger.parser.OpenAPIParser
import io.swagger.v3.parser.core.models.SwaggerParseResult
import io.swagger.v3.oas.models.OpenAPI
import io.swagger.v3.oas.models.media.Schema
import java.io.File

// This script reads a Swagger 2.0 specification from a file and converts it to OpenAPI 3.0.
// It also ensures that all schemas starting with "Nullable" have the nullable=true property.

fun main() {
  val inputFilename = args.firstOrNull() ?: error("Please provide a path to input specification file.")

  val inputText = File(inputFilename).readText()

  val output: SwaggerParseResult = OpenAPIParser().readContents(inputText, null, null)
    ?: throw RuntimeException("Failed to parse input file");

  if (output.openAPI == null) {
    throw RuntimeException("No OpenAPI object found")
  }

  // Post-process the schema to add nullable property for schemas with Nullable prefix
  processNullableSchemas(output.openAPI)

  io.swagger.v3.core.util.Json.prettyPrint(output.openAPI)
}

fun processNullableSchemas(openAPI: OpenAPI) {
  val schemas = openAPI.components?.schemas ?: return
  
  schemas.forEach { (name, schema) ->
    if (name.startsWith("Nullable")) {
      // Make sure the schema is marked as nullable
      schema.nullable = true
    }
  }
}

main()

/*
*  Copyright (c) WSO2 Inc. (http://www.wso2.org) All Rights Reserved.
*
*  WSO2 Inc. licenses this file to you under the Apache License,
*  Version 2.0 (the "License"); you may not use this file except
*  in compliance with the License.
*  You may obtain a copy of the License at
*
*    http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing,
* software distributed under the License is distributed on an
* "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
* KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations
* under the License.
 */

package testutils

const CustomTestExportDirectory = "CustomExportDirectory"
const TestSwagger2DefinitionPath = "testdata/swagger2Definition.yaml"
const TestOpenAPI3DefinitionPath = "testdata/openAPI3Definition.yaml"
const TestOpenAPISpecificationURL = "https://petstore.swagger.io/v2/swagger.json"
const TestMigrationDirectorySuffix = "/migration"

const DefaultApictlTestAppName = "default-apictl-app"

//Export test cases
const DevFirstDefaultAPIName = "SwaggerPetstoreNew"
const DevFirstSwagger2APIName = "PizzaShackAPI"
const DevFirstDefaultAPIVersion = "1.0.0"
const APIYamlFilePath = "/api.yaml"
const SampleAPIYamlFilePath = "testdata/sample-api.yaml"
const TestArtifact1Path = "testdata/TestArtifactDirectory/ArtifactSet1"
const TestCase1DocName = "/Docs/Doc1"
const TestCase1DocPath = TestArtifact1Path + "/Doc1/testDoc.pdf"
const TestCase1DestPathSuffix = TestCase1DocName + "/testDoc.pdf"
const TestCase1DocMetaDataPath = TestArtifact1Path + "/Doc1/document.yaml"
const TestCase1DestMetaDataPathSuffix = TestCase1DocName + "/document.yaml"
const TestArtifact2Path = "testdata/TestArtifactDirectory/ArtifactSet2"
const TestCase2DocName = "/Docs/Doc2"
const TestCase2DocPath = TestArtifact2Path + "/Doc2/mockPdf.pdf"
const TestCase2DestPathSuffix = TestCase2DocName + "/mockPdf.pdf"
const TestCase2DestPathSuffixForUpdate = TestCase2DocName + "/testDoc.pdf"
const TestCase2DocMetaDataPath = TestArtifact2Path + "/Doc2/document.yaml"
const TestCase2DestMetaDataPathSuffix = TestCase2DocName + "/document.yaml"
const TestCase2PngPath = TestArtifact2Path + "/icon.png"
const TestCase2DestPngPathSuffix = "/Image/icon.png"
const TestCase2JpegPath = TestArtifact2Path + "/icon.jpeg"
const TestCase2DestJpegPathSuffix = "/Image/icon.jpeg"
const TestDefaultExtractedFileName = "/SwaggerPetstoreNew-1.0.0"

// APIEndpointParamsFile : Endpoint URL api_params.yaml
const APIEndpointParamsFile = "testdata/api_params_endpoint.yaml"

// APIEndpointRetryTimeoutParamsFile : Endpoint URL and Retry Timeout api_params.yaml
const APIEndpointRetryTimeoutParamsFile = "testdata/api_params_endpoint_retrytimeout.yaml"

// APISecurityFalseParamsFile : Security false api_params.yaml
const APISecurityFalseParamsFile = "testdata/api_params_security_false.yaml"

// APISecurityDigestParamsFile : Security Digest api_params.yaml
const APISecurityDigestParamsFile = "testdata/api_params_security_digest.yaml"

// APISecurityBasicParamsFile : Security Basic api_params.yaml
const APISecurityBasicParamsFile = "testdata/api_params_security_basic.yaml"

// APIFullParamsFile : Full api_params.yaml
const APIFullParamsFile = "testdata/api_params_full.yaml"

// APIProductFullParamsFile : Full api_product_params.yaml
const APIProductFullParamsFile = "testdata/api_product_params_full.yaml"

// CertificatesDirectoryPath : Directory path for the dunnmy certificates
const CertificatesDirectoryPath = "testdata/TestArtifactDirectory/certificates"

// UnlimitedPolicy : Unlimited Throttle Policy
const UnlimitedPolicy = "Unlimited"

// API types
const APITypeREST = "HTTP"
const APITypeSoap = "SOAP"
const APITypeSoapToRest = "SOAPTOREST"
const APITypeGraphQL = "GraphQL"
const APITypeWebScoket = "WS"
const APITypeWebSub = "WEBSUB"
const APITypeSSE = "SSE"

// REST API Endpoint URL
const RESTAPIEndpoint = "petstore.swagger.io"

// SOAP API Endpoint URL
const SoapEndpointURL = "http://ws.cdyne.com/phoneverify/phoneverify.asmx"

// GraphQL API Endpoint URL
const GraphQLEndpoint = "http://www.mocky.io/v2/5ea84def2d0000a52d3a3ecd"

// Web Socket API Endpoint URL
const WebSocketEndpoint = "ws://echo.websocket.org:80"

// APIHttpRestEndpointWithoutLoadBalancingOrFailoverParamsFile : HTTP/REST Endpoint without Loadbalancing and Failover URLs in api_params.yaml
const APIHttpRestEndpointWithoutLoadBalancingOrFailoverParamsFile = "testdata/api_params_http_rest_endpoint_without_lb_or_failover.yaml"

// APIHttpSoapEndpointWithoutLoadBalancingOrFailoverParamsFile : HTTP/SOAP Endpoint without Loadbalancing and Failover URLs in api_params.yaml
const APIHttpSoapEndpointWithoutLoadBalancingOrFailoverParamsFile = "testdata/api_params_http_soap_endpoint_without_lb_or_failover.yaml"

// APIHttpRestEndpointWithLoadBalancingParamsFile : HTTP/REST Endpoint with Loadbalancing URLs in api_params.yaml
const APIHttpRestEndpointWithLoadBalancingParamsFile = "testdata/api_params_http_rest_endpoint_with_load_balancing.yaml"
package servicemap

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"github.com/Azure/go-autorest/tracing"
	"net/http"
)

// MapsClient is the service Map API Reference
type MapsClient struct {
	BaseClient
}

// NewMapsClient creates an instance of the MapsClient client.
func NewMapsClient(subscriptionID string) MapsClient {
	return NewMapsClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewMapsClientWithBaseURI creates an instance of the MapsClient client using a custom endpoint.  Use this when
// interacting with an Azure cloud that uses a non-standard base URI (sovereign clouds, Azure stack).
func NewMapsClientWithBaseURI(baseURI string, subscriptionID string) MapsClient {
	return MapsClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// Generate generates the specified map.
// Parameters:
// resourceGroupName - resource group name within the specified subscriptionId.
// workspaceName - OMS workspace containing the resources of interest.
// request - request options.
func (client MapsClient) Generate(ctx context.Context, resourceGroupName string, workspaceName string, request BasicMapRequest) (result MapResponse, err error) {
	if tracing.IsEnabled() {
		ctx = tracing.StartSpan(ctx, fqdn+"/MapsClient.Generate")
		defer func() {
			sc := -1
			if result.Response.Response != nil {
				sc = result.Response.Response.StatusCode
			}
			tracing.EndSpan(ctx, sc, err)
		}()
	}
	if err := validation.Validate([]validation.Validation{
		{TargetValue: resourceGroupName,
			Constraints: []validation.Constraint{{Target: "resourceGroupName", Name: validation.MaxLength, Rule: 64, Chain: nil},
				{Target: "resourceGroupName", Name: validation.MinLength, Rule: 1, Chain: nil},
				{Target: "resourceGroupName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_-]+`, Chain: nil}}},
		{TargetValue: workspaceName,
			Constraints: []validation.Constraint{{Target: "workspaceName", Name: validation.MaxLength, Rule: 63, Chain: nil},
				{Target: "workspaceName", Name: validation.MinLength, Rule: 3, Chain: nil},
				{Target: "workspaceName", Name: validation.Pattern, Rule: `[a-zA-Z0-9_][a-zA-Z0-9_-]+[a-zA-Z0-9_]`, Chain: nil}}}}); err != nil {
		return result, validation.NewError("servicemap.MapsClient", "Generate", err.Error())
	}

	req, err := client.GeneratePreparer(ctx, resourceGroupName, workspaceName, request)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MapsClient", "Generate", nil, "Failure preparing request")
		return
	}

	resp, err := client.GenerateSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "servicemap.MapsClient", "Generate", resp, "Failure sending request")
		return
	}

	result, err = client.GenerateResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "servicemap.MapsClient", "Generate", resp, "Failure responding to request")
	}

	return
}

// GeneratePreparer prepares the Generate request.
func (client MapsClient) GeneratePreparer(ctx context.Context, resourceGroupName string, workspaceName string, request BasicMapRequest) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
		"workspaceName":     autorest.Encode("path", workspaceName),
	}

	const APIVersion = "2015-11-01-preview"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsPost(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.OperationalInsights/workspaces/{workspaceName}/features/serviceMap/generateMap", pathParameters),
		autorest.WithJSON(request),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GenerateSender sends the Generate request. The method will close the
// http.Response Body if it receives an error.
func (client MapsClient) GenerateSender(req *http.Request) (*http.Response, error) {
	return client.Send(req, azure.DoRetryWithRegistration(client.Client))
}

// GenerateResponder handles the response to the Generate request. The method always
// closes the http.Response Body.
func (client MapsClient) GenerateResponder(resp *http.Response) (result MapResponse, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

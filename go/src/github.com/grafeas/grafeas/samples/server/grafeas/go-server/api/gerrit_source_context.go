/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * OpenAPI spec version: 0.1
 *
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package swagger

// A SourceContext referring to a Gerrit project.
type GerritSourceContext struct {

	// The URI of a running Gerrit instance.
	HostUri string `json:"hostUri,omitempty"`

	// The full project name within the host. Projects may be nested, so \"project/subproject\" is a valid project name. The \"repo name\" is hostURI/project.
	GerritProject string `json:"gerritProject,omitempty"`

	// A revision (commit) ID.
	RevisionId string `json:"revisionId,omitempty"`

	// The name of an alias (branch, tag, etc.).
	AliasName string `json:"aliasName,omitempty"`

	// An alias, which may be a branch or tag.
	AliasContext AliasContext `json:"aliasContext,omitempty"`
}

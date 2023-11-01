/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package snapshots

import (
	"github.com/containerd/containerd/v2/metadata"
	"github.com/containerd/containerd/v2/plugin"
	"github.com/containerd/containerd/v2/plugin/registry"
	"github.com/containerd/containerd/v2/plugins"
	"github.com/containerd/containerd/v2/services"
)

func init() {
	registry.Register(&plugin.Registration{
		Type: plugins.ServicePlugin,
		ID:   services.SnapshotsService,
		Requires: []plugin.Type{
			plugins.MetadataPlugin,
		},
		InitFn: func(ic *plugin.InitContext) (interface{}, error) {
			m, err := ic.Get(plugins.MetadataPlugin)
			if err != nil {
				return nil, err
			}

			return m.(*metadata.DB).Snapshotters(), nil
		},
	})
}

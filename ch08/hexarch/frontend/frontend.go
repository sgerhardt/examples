/*
 * Copyright 2023 Matthew A. Titmus
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

package frontend

import (
	"github.com/cloud-native-go/examples/ch08/hexarch/core"
)

type FrontEnd interface {
	Start(kv *core.KeyValueStore) error
}

type zeroFrontEnd struct{}

func (f zeroFrontEnd) Start(kv *core.KeyValueStore) error {
	return nil
}

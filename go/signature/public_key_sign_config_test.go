// Copyright 2017 Google Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
////////////////////////////////////////////////////////////////////////////////
package signature_test

import (
  "testing"
  "github.com/google/tink/go/tink/tink"
  "github.com/google/tink/go/signature/signature"
)

func TestPublicKeySignConfigInstance(t *testing.T) {
  config := signature.PublicKeySignConfig()
  if config == nil {
    t.Errorf("instance of publicKeySignConfig is nil")
  }
}

func TestPublicKeySignConfigRegistration(t *testing.T) {
  _, err := signature.PublicKeySignConfig().RegisterStandardKeyTypes()
  if err != nil {
    t.Errorf("cannot register standard key types")
  }
  // check for EcdsaSignKeyManager
  keyManager, err := tink.Registry().GetKeyManager(signature.ECDSA_SIGN_TYPE_URL)
  if err != nil {
    t.Errorf("unexpected error: %s", err)
  }
  var _ = keyManager.(*signature.EcdsaSignKeyManager)
}
/*
 * Copyright 2012-2019 the original author or authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package gs_test

import (
	"os"
	"testing"
	"time"

	"github.com/go-spring/spring-core/assert"
	"github.com/go-spring/spring-core/gs"
	"github.com/go-spring/spring-core/util"
)

func startApplication(cfgLocation ...string) (*gs.App, gs.Pandora) {
	app := gs.NewApp()
	app.EnablePandora()

	var p gs.Pandora
	app.Config(func(b gs.Pandora) { p = b })

	go app.Run(cfgLocation...)
	time.Sleep(100 * time.Millisecond)
	return app, p
}

func TestConfig(t *testing.T) {

	t.Run("config via env", func(t *testing.T) {
		os.Clearenv()
		_ = os.Setenv(util.SpringProfile, "dev")
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(util.SpringProfile), "dev")
	})

	t.Run("config via env 2", func(t *testing.T) {
		os.Clearenv()
		_ = os.Setenv(gs.SPRING_PROFILE, "dev")
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(util.SpringProfile), "dev")
	})

	t.Run("profile via config", func(t *testing.T) {
		os.Clearenv()
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(util.SpringProfile), "test")
	})

	t.Run("profile via env&config", func(t *testing.T) {
		os.Clearenv()
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(util.SpringProfile), "test")
	})

	t.Run("profile via env&config 2", func(t *testing.T) {
		os.Clearenv()
		_ = os.Setenv(gs.SPRING_PROFILE, "dev")
		app, p := startApplication("testdata/config/")
		defer app.ShutDown()
		assert.Equal(t, p.Prop(util.SpringProfile), "dev")
	})
}

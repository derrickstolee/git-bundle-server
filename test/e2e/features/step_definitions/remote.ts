import { Given, } from '@cucumber/cucumber'
import { RemoteRepo } from '../classes/remote'
import { BundleServerWorld } from '../support/world'
import * as path from 'path'
import { env } from 'process'

/**
 * Steps relating to the setup of the remote repository users will clone from.
 */

Given('a remote repository {string}', async function (this: BundleServerWorld, url: string) {
  this.remote = new RemoteRepo(false, url)
})

Given('a bundle URI from TEST_BUNDLE_SERVER_URI', async function (this: BundleServerWorld) {
  this.bundleURIBase = env["TEST_BUNDLE_SERVER_URI"]

  if (this.bundleURIBase == "") {
    throw new Error("This test requires TEST_BUNDLE_SERVER_URI")
  }
})

Given('a bundle route {string}', async function (this: BundleServerWorld, route: string) {
  this.bundleURI = this.bundleURIBase + route
})

Given('a new remote repository with main branch {string}', async function (this: BundleServerWorld, mainBranch: string) {
  this.remote = new RemoteRepo(true, path.join(this.trashDirectory, "server"), mainBranch)
})

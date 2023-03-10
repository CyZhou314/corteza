package app

import (
	"github.com/cyzhou314/corteza/server/codegen/schema"
	"github.com/cyzhou314/corteza/server/app/options"
	"github.com/cyzhou314/corteza/server/system"
	"github.com/cyzhou314/corteza/server/compose"
	"github.com/cyzhou314/corteza/server/automation"
	"github.com/cyzhou314/corteza/server/federation"
)

corteza: schema.#platform & {
	"ident": "corteza"

	"options": [
		options.DB,
		options.HTTPClient,
		options.HTTPServer,
		options.RBAC,
		options.SCIM,
		options.SMTP,
		options.actionLog,
		options.apigw,
		options.auth,
		options.corredor,
		options.environment,
		options.eventbus,
		options.federation,
		options.limit,
		options.locale,
		options.log,
		options.messagebus,
		options.monitor,
		options.objectStore,
		options.provision,
		options.sentry,
		options.template,
		options.upgrade,
		options.waitFor,
		options.websocket,
		options.workflow,
		options.discovery,
	]

	// platform resources
	"resources": resources

	"components": [
		system.component,
		compose.component,
		automation.component,
		federation.component,
	]
}

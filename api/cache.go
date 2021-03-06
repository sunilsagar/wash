package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/puppetlabs/wash/journal"
	"github.com/puppetlabs/wash/plugin"
)

// swagger:route DELETE /cache cache cacheDelete
//
// Remove items from the cache
//
// Removes the specified entry and its children from the cache.
//
//     Produces:
//     - application/json
//
//     Schemes: http
//
//     Responses:
//       200:
//       500: errorResp
var cacheHandler handler = func(w http.ResponseWriter, r *http.Request, p params) *errorResponse {
	path := p.Path
	ctx := r.Context()
	journal.Record(ctx, "API: Cache DELETE %v", path)
	deleted, err := plugin.ClearCacheFor(path)
	if err != nil {
		journal.Record(ctx, "API: Cache DELETE flush cache errored constructing regexp from %v: %v", path, err)
		return unknownErrorResponse(fmt.Errorf("Could not use path %v in a regexp: %v", path, err))
	}

	w.WriteHeader(http.StatusOK)
	jsonEncoder := json.NewEncoder(w)
	if err = jsonEncoder.Encode(deleted); err != nil {
		journal.Record(ctx, "API: Cache DELETE marshalling %v errored: %v", path, err)
		return unknownErrorResponse(fmt.Errorf("Could not marshal deleted keys for %v: %v", path, err))
	}

	journal.Record(ctx, "API: Cache DELETE %v complete: %v", path, deleted)
	return nil
}

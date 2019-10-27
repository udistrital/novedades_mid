package models

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/astaxie/beego"
)

type JbpmReplica struct {
	NovedadReplicaCollection struct {
		novedad_replica []struct {
			id                int    `json:"id"`
			ArgonovedadId     int    `json:"argonovedad_id"`
			NovedadId         int    `json:"novedad_id"`
			Activo            bool   `json:"activo"`
			FechaCreacion     string `json:"fecha_creacion"`
			FechaModificacion string `json:"fecha_modificacion"`
		} `json:"novedad_replica"`
	} `json:"novedad_replicaCollection"`
}

func GetJsonWSO2(urlp string, target interface{}) error {
	b := new(bytes.Buffer)
	client := &http.Client{}
	req, err := http.NewRequest("GET", urlp, b)
	req.Header.Set("Accept", "application/json")
	r, err := client.Do(req)
	if err != nil {
		beego.Error("error", err)
		return err
	}
	defer func() {
		if err := r.Body.Close(); err != nil {
			beego.Error(err)
		}
	}()

	return json.NewDecoder(r.Body).Decode(target)
}

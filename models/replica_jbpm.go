package models

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

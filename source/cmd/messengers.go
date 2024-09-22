package main

// func MakeMessengers(dataDB, gatewayDB *sql.DB, dataSchema, gatewaySchema string, worker *worker.Worker) (messenger.Messengers, error) {
// 	Lembaga := table.Lembaga.FromSchema(dataSchema)
// 	stmt := Lembaga.SELECT(Lembaga.Memberid, Lembaga.Nama).LIMIT(3)

// 	var model []model.Lembaga
// 	if err := stmt.Query(dataDB, &model); err != nil {
// 		return nil, err
// 	}

// 	worker.Run()

// 	messengers := make([]messenger.Messenger, 0)
// 	for _, m := range model {
// 		m := m
// 		messengers = append(messengers, messenger.Messenger{
// 			TenantID:      m.Memberid,
// 			TenantName:    *m.Nama,
// 			DataSchema:    dataSchema,
// 			GatewaySchema: gatewaySchema,
// 			DataDB:        dataDB,
// 			GatewayDB:     gatewayDB,
// 			Worker:        worker,
// 			Provider:      httpprovider.NewHTTPProvider(),
// 		})
// 	}

// 	return messengers, nil
// }

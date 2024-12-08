package main

import (
	"context"
	"fmt"
	"log"
)

type ChairA struct {
	ChairID string
	Dist    int
}

func getChairsSortedByCurrentPosition(
	ctx context.Context,
	curLat int,
	curLng int,
) ([]ChairA, error) {
	chairLocs := []ChairA{}
	rows, err := db.QueryxContext(ctx, fmt.Sprintf(`
select
    chair_id as chair_id,
    abs(tbl2.latitude - (%d)) + abs(tbl2.longitude - (%d)) as dist
from
    (
        select
            chair_locations.chair_id as chair_id,
            chair_locations.latitude as latitude,
            chair_locations.longitude as longitude
        from
            chair_locations
            inner join
                (
                    select
                        chair_id,
                        max(created_at) as max_created_at
                    from
                        chair_locations
                    group by
                        chair_id
                ) as tbl1
            on  tbl1.chair_id = chair_locations.chair_id
            AND tbl1.max_created_at = chair_locations.created_at
    ) tbl2
    inner join
        chairs
    on  chairs.id = tbl2.chair_id AND chairs.is_active = TRUE
order by
    dist
;
	`, curLat, curLng))
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		c := ChairA{}
		err := rows.Scan(&c.ChairID, &c.Dist)
		if err != nil {
			log.Fatalln(err)
		}
		chairLocs = append(chairLocs, c)
	}

	return chairLocs, nil
}

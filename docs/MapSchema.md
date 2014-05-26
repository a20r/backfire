Map Schema
==========

Documentation about the standard map description schema in JSON.

## Overall Map Structure

	{
		"name": <String -- Name of map>,
		"author_name": <String -- Name of author>,
		"author_email": <String -- Email of author>,
		"map": {
			"width": <Int -- Width of map>,
			"height": <Int -- Height of map>,
			"depth": <Int -- Optional -- 3D depth of map>,
			"configurations": [
				{
					"initial": {
						"x": <Int -- X coordinate of initial configuration>,
						"y": <Int -- Y coordinate of initial configuration>,
						"z": <Int -- Optional -- Z coordinate of initial configuration>
					},
					"goal": {
						"x": <Int -- X coordinate of goal configuration>,
						"y": <Int -- Y coordinate of goal configuration>,
						"z": <Int -- Optional -- Z coordinate of goal configuration>
				}, …
			]
			"obstacles": [
				{
					"dynamic": <Boolean -- Whether obstacle is dynamic>,
					"name": <String -- Optional -- Name of obstacle>,
					"vertices": [
						{
							"x": <Int -- X coordinate of vertex>,
							"y": <Int-- Y coordinate of vertex>,
							"z": <Int -- Optional -- Z coordinate of vertex>
						}, …
					]
				}, …
			] 	
		}
	}
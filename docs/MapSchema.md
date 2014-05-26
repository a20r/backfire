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
			"obstacles": [
				{
					"dynamic": <Boolean -- Whether obstacle is dynamic>,
					"name": <String -- Optional -- Name of obstacle>,
					"vertices": [
						{
							"x": <Int -- X coordinate of vertex>,
							"y": <Int-- Y coordinate of vertex>,
							"z": <Int -- Optional -- Z coordinate of vertex>
						}, ...
					]
				}, ...
			] 	
		}
	}
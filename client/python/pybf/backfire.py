
import urllib2
import urllib
import json
import scene


class Backfire(object):

    def __init__(self, **kwargs):
        self.host = kwargs.get("host", "localhost")
        self.port = kwargs.get("port", 8000)
        self.base_url = "http://" + self.host + ":" + str(self.port)

    def get_maps_by_author(self, name):
        name = urllib.quote(name)
        req = urllib2.urlopen(self.base_url + "/maps/authorName/" + name)
        res_str = req.read()
        res_list = json.loads(res_str)

        scene_list = list()

        for scene_json in res_list:
            scene_list.append(scene.Scene(**scene_json))

        return scene_list

    def get_map_by_name(self, name):
        name = urllib.quote(name)
        req = urllib2.urlopen(self.base_url + "/maps/name/" + name)
        res_str = req.read()
        res_dict = json.loads(res_str)

        sc = scene.Scene(**res_dict)
        return sc





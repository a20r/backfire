
import point


class Obstacle(object):

    def __init__(self, **kwargs):
        self.dynamic = kwargs.get("dynamic")
        self.name = kwargs.get("name", None)
        self.vertices = self.init_vertices(kwargs.get("vertices"))

    def init_vertices(self, v_list):
        vertices = list()
        for v_dict in v_list:
            vertices.append(point.Point(
                v_dict.get("x"),
                v_dict.get("y"),
                v_dict.get("z", 0)
            ))

        return vertices

    def is_dynamic(self):
        return self.dynamic

    def get_name(self):
        return self.name

    def get_vertices(self):
        return self.vertices



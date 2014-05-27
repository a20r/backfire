
import configuration
import obstacle


class Scene(object):

    def __init__(self, **kwargs):
        self.name = kwargs.get("name")
        self.author_name = kwargs.get("authorName")
        self.author_email = kwargs.get("authorEmail")
        self.width = kwargs.get("width")
        self.height = kwargs.get("height")
        self.depth = kwargs.get("depth", 0)
        self.configurations = self.init_configurations(
            kwargs.get("configurations")
        )
        self.obstacles = self.init_obstacles(
            kwargs.get("obstacles")
        )

    def init_configurations(self, configs_json):
        configs = list()
        for config in configs_json:
            configs.append(configuration.Configuration(**config))

        return configs

    def init_obstacles(self, obs_json):
        obstacles = list()
        for ob in obs_json:
            obstacles.append(obstacle.Obstacle(**ob))

        return obstacles

    def get_name(self):
        return self.name

    def get_author_name(self):
        return self.author_name

    def get_author_email(self):
        return self.author_email

    def get_width(self):
        return self.width

    def get_height(self):
        return self.height

    def get_depth(self):
        return self.depth

    def get_configurations(self):
        return self.configurations

    def get_obstacles(self):
        return self.obstacles


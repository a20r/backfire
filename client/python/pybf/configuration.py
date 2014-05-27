
import point


class Configuration(object):

    def __init__(self, **kwargs):
        self.initial = kwargs.get("initial")
        self.goal = kwargs.get("goal")

    def get_initial(self):
        return self.initial

    def get_goal(self):
        return self.goal



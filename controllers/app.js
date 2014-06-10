function AppController(){}

AppController.prototype.register = function register(router) {
	router.get('/', this.root);
	router.get('/about', this.about);
};

AppController.prototype.root = function root(req, res) {
	res.render('index');
};

AppController.prototype.about = function about(req, res) {
	res.render('about');
};

module.exports = AppController;
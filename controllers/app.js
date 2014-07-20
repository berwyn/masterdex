function AppController(){}

AppController.prototype.register = function register(router) {
	router.get('/', this.root);
	router.get('/about', this.about);
};

AppController.prototype.root = function root(req, res) {
	res.render('app/index');
};

AppController.prototype.about = function about(req, res) {
	res.render('app/about');
};

module.exports = AppController;
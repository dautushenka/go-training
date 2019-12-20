'use strict';

(function(history){
    var pushState = history.pushState;
    history.pushState = function(state, title, url) {
        if (typeof history.onpushstate == "function") {
            history.onpushstate({state, title, url});
        }
        // ... whatever else you want to do
        // maybe call onhashchange e.handler
        return pushState.apply(history, arguments);
    };
})(window.history);

window.history.onpushstate = function(e) {
    render(e.url, e.state);
};

function render(pathname, state) {
    const user = window.user;
    let Render = Articles;
    let props = {};
    const path = pathname.split('/');
    if (path[1] === 'post') {
        Render = Post;
        props.post = state.post;
        props.user = user;
    } else if (path[1] === 'category') {
        props.category = parseInt(path[2]);
    } else if (path[1] === 'login') {
        Render = LoginPage
    }

    ReactDOM.render(
        React.createElement(Render, props, null), document.getElementById('content')
    );

    ReactDOM.render(
        React.createElement(Categories, {selected: props.category}, null), document.getElementById('categories')
    );

    ReactDOM.render(
        React.createElement(HeaderList, {user}, null), document.getElementById('header-list')
    );
}

document.addEventListener("DOMContentLoaded", () => {
    fetch("/api/v1/auth", {
        headers: {
            Authentication: 'Bearer ' + (localStorage.getItem('token') || ''),
        }
    }).then(response => {
        response.ok && response.json().then(data => {
            window.user = data.user;
            render(location.pathname, history.state)
        })
    });

    render(location.pathname)
});

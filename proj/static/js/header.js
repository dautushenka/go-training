'use strict';

function HeaderList({user}) {
    const onClick = (e) => {
        e.preventDefault();
        history.pushState({}, null, '/login')
    };

    const HPClick = (e) => {
        e.preventDefault();
        history.pushState({}, null, '/')
    };

    const onLogout = (e) => {
        e.preventDefault();
        window.user = undefined;
        window.localStorage.removeItem('token');
        window.history.pushState({}, '', '/');
    };

    return [
        React.createElement('li', {key: 'hp'},
            React.createElement('a', {href: "/", onClick: HPClick}, 'Home')
        ),
        React.createElement('li', {key: 'user'},
            user ?
                [
                    React.createElement('div', null, `Hello, ${user.first_name}`),
                    React.createElement('a', {onClick: onLogout, href: "#", className: 'logout'}, "Logout")
                ]
                : React.createElement('a', {href: "/login", onClick}, "Login")
        ),
    ]
}
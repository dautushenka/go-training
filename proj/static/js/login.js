'use strict';

class LoginPage extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {
            login: '',
            password: '',
        }
    }

    onLoginUpdate = (e) => this.setState({login: e.target.value});
    onPasswordUpdate = (e) => this.setState({password: e.target.value});

    onSubmit = (e) => {
        e.preventDefault();
        fetch('/api/v1/auth', {
            method: 'POST',
            body: JSON.stringify(this.state)
        }).then(response => {
            if (response.ok) {
                response.json().then(data => {
                    window.localStorage.setItem('token', data.token);
                    window.user = data.user;
                    window.history.pushState({}, '', '/');
                })
            } else {
                response.json().then(data => {
                    alert(data.message)
                })
            }
        })
    };

    render() {
        return React.createElement('div', {className: 'login-form'},
            React.createElement('form', {action: "", onSubmit: this.onSubmit}, [
                React.createElement('div', null, 'Login/password is user/user'),
                React.createElement('input', {
                    name: "login",
                    placeholder: "Login",
                    onChange: this.onLoginUpdate,
                    value: this.state.login,
                }, null),
                React.createElement('input', {
                    name: "password",
                    placeholder: "Password",
                    type: 'password',
                    onChange: this.onPasswordUpdate,
                    value: this.state.password,
                }, null),
                React.createElement('input', {type: 'submit'}, null),
            ])
        )
    }
}
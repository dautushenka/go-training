'use strict';


function Category({id, name, isSelected}) {
    return React.createElement(
        'li',
        {
            className: isSelected ? 'selected' : undefined,
        },
        React.createElement('a', {
            href: `/category/${id}`,
            onClick: (e) => {e.preventDefault(); history.pushState({}, name, `/category/${id}`)}
        }, name)
    )
}

class Categories extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {categories: []}
    }

    componentDidMount() {
        fetch("/api/v1/category").then(response => {
            response.json().then(categories => {
                this.setState({categories})
            })
        })
    }

    render() {
        return React.createElement(
            'ul',
            null,
            this.state.categories.map(cat => {
                return React.createElement(Category, {key: cat.id, ...cat, isSelected: cat.id === this.props.selected}, null);
            })
        );
    }
}
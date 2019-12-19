'use strict';

function ArticleTitle({title, id}) {
    return React.createElement(
        'h3',
        null,
        React.createElement(
            'a',
            {
                onClick: (e) => {
                    e.preventDefault();
                    console.log('click')
                },
                href: `/post/${id}`
            }
            , `${title}`)
    );
}

function ArticleBody({body}) {
    return React.createElement('p', null, `${body}`);
}

function Article(post) {
    return React.createElement('article', null, [
        React.createElement(ArticleTitle, post, null),
        React.createElement(ArticleBody, post, null)
    ]);
}

class Articles extends React.PureComponent {
    constructor(props) {
        super(props);
        this.state = {posts: []}
    }

    componentDidMount() {
        fetch("/api/v1/post").then(response => {
            response.json().then(posts => {
                this.setState({posts})
            })
        })
    }

    render() {
        return this.state.posts.map(post => {
            return React.createElement(Article, {key: post.id, ...post}, null);
        })
    }
}

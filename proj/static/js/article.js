'use strict';

function ArticleTitle(post) {
    const  {title, id} = post;
    return React.createElement(
        'h3',
        null,
        React.createElement(
            'a',
            {
                onClick: (e) => {
                    e.preventDefault();
                    window.history.pushState({post}, title, `/post/${id}`);
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
        this.fetchPosts()
    }

    componentDidUpdate(prevProps) {
        if (prevProps.category !== this.props.category) {
            this.fetchPosts()
        }
    }

    fetchPosts = () => {
        let URL = "/api/v1/post";
        if (this.props.category) {
            URL += "?category=" + encodeURIComponent(this.props.category);
        }
        fetch(URL).then(response => {
            response.json().then(posts => {
                this.setState({posts})
            })
        })
    };

    render() {
        return this.state.posts.map(post => {
            return React.createElement(Article, {key: post.id, ...post}, null);
        })
    }
}

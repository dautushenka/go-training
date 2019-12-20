'use strict';

function PostActions({post, user}) {
    if (!user || post.user_id !== user.id) {
        return null;
    }

    const onDelete = (e) => {
        e.preventDefault();
        if (window.confirm("Are you sure?")) {
            fetch(`/api/v1/post/${post.id}`, {
                method: 'DELETE',
                headers: {
                    Authentication: 'Bearer ' + (localStorage.getItem('token') || ''),
                }
            }).then(() => history.pushState({}, '', '/'))
        }
    };

    const onEdit = (e) => {
        e.preventDefault();

        // Open form
    };

    return React.createElement('div', { className: 'postActions' }, [
        React.createElement('a', {key: 'edit', href: `/post/${post.id}/edit`, onClick: onEdit}, "Edit"),
        React.createElement('a', {key: 'delete', href: `#`, onClick: onDelete}, "Delete")
    ])
}

function Post({post, user}) {
    const {title, body} = post;
    return React.createElement('article', null, [
        React.createElement('h1', null, title),
        React.createElement('div', null, body),
        React.createElement(PostActions, {post, user}, null),
    ])
}
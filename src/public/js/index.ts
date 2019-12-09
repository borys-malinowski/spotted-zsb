const allPosts: HTMLElement | null = document.querySelector("#all-posts");

async function getPosts() {
    const response: Response = await fetch(location.origin + "/api/get-posts");
    const posts = await response.json();
    const allPostsArray: HTMLDivElement[] = [];
    for(let i: number = 0; i < posts.length; i++) {
        if(i % 2 === 0) {
            const parsedPost = JSON.parse(posts[i]);
            const el: HTMLDivElement = document.createElement("div");
            el.classList.add("post");
            el.innerHTML = `<h2>${parsedPost.title}</h2><br>${parsedPost.content}`;
            allPostsArray.unshift(el);
        }
    }
    for(let j: number = 0; j < allPostsArray.length; j++) {
        allPosts!.appendChild(allPostsArray[j]);
    }
}

getPosts();

import  API  from "../services/API.js";

export class HomePage extends HTMLElement{
      
  async render(){
    const topMovies = await API.getTopMovies();
      function renderMoviesInList(movies, ul){
      ul.innerHTML = ''; // Clear existing content
      movies.forEach(movie => {
        const li = document.createElement("li");
        li.textContent = movie.title;
        ul.appendChild(li);
      });
    }
    renderMoviesInList(topMovies, document.getElementById("top10-list"));



  

  }
  
  
  connectedCallback(){
        const template = document.getElementById("home-template");
        const clone = template.content.cloneNode(true);
        this.appendChild(clone);

        this.render();
      }


}

customElements.define("home-page", HomePage);
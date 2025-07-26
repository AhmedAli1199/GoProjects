import { HomePage } from "./components/HomePage.js";

window.addEventListener('DOMContentLoaded', function() {
document.querySelector("main").appendChild(new HomePage());

}
);

window.app = {

    search: function(event){
        event.preventDefault();
        const q = this.document.querySelector("#search").value;
    }

}
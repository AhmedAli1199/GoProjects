export const API = {
    baseUrl: "api",

    getTopMovies: async () => {
        return await API.fetch('movies/top');
    },
    getRandomMovies: async () => {
        return await API.fetch('movies/random');
    },
    serachMovies: async (query, order, genre)=> {
        const args = {};
        if (query) args.query = query;
        if (order) args.order = order;
        if (genre) args.genre = genre;

        return await API.fetch('movies/search', args);
    },
    getMovieById: async (id) => { 
        if (!id) throw new Error("Movie ID is required");
        return await API.fetch(`movies/${id}`);
    },
    fetch: async (serviceUrl, args) => {

        try{
        const queryString = args ? new URLSearchParams(args).toString() : '';  
        console.log(`Fetching from: ${API.baseUrl}/${serviceUrl}?${queryString}`);
        const response = await fetch(`${API.baseUrl}/${serviceUrl}?${queryString}`)
        const data = await response.json();
        if (!response.ok) {
            throw new Error(`Error: ${data.message || 'Something went wrong'}`);
        }
        return data;
        } catch (error) {
        console.error("API Error:", error);
        }

        
    }



}

export default API;
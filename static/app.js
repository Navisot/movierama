const vm = new Vue({
    delimiters: ['{(', ')}'],
    el: '#app',
    data: {
        isGallery: false,
        movies:[],
        movie:{
            title:'',
            description:''
        },
    },
    methods: {
        saveMovie() {
            var that = this;
            axios.post("/api/save/movie", {"title": this.movie.title, "description": this.movie.description}).then(function(response){
               if (response.data.status == "OK"){
                   that.movie.title = '';
                   that.movie.description = '';
                    alert('Success!')
               } else {
                   alert ('Fail');
               }
            });
        },
        showMoviesOrderBy(item) {
            axios.get("/api/show/movies/sort/"+item).then(response => {
                this.movies = response.data;
            });
        },
        showUserMovies(user_id) {
            vm.isGallery = true;
            // show movies where user_id = user_id

        }
    },
    mounted: function(){
        axios.get("/api/show/movies").then(response => {
            this.movies = response.data;
        });
    }
});
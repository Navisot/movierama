const vm = new Vue({
    delimiters: ['{(', ')}'],
    el: '#app',
    data: {
        isLoggedIn: false,
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
        },
        likeThisMovie(movie_id) {
            if (vm.isLoggedIn) {
                axios.post('/api/rate/movie/'+movie_id+'/1').then(response => {
                    console.log(response.data)
                });
            } else {
                alert('Please login in order to like or hate a movie!')
            }

        },
        hateThisMovie(movie_id) {
            if (vm.isLoggedIn) {
                axios.post('/api/rate/movie/'+movie_id+'/2').then(response => {
                    console.log(response.data)
                });
            } else {
                alert('Please login in order to like or hate a movie!')
            }
        }
    },
    mounted: function(){
        axios.get("/api/show/movies").then(response => {
            this.movies = response.data;
        });
        axios.get("/api/user/status").then(response => {
            this.isLoggedIn = response.data.is_logged_in;
        });
    }
});
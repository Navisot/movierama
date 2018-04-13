const vm = new Vue({
    delimiters: ['{(', ')}'],
    el: '#app',
    data: {
        selectedUser: '',
        selectedUserID: 0,
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
        showMoviesOrderBy(item, movie_user_id) {
                axios.get("/api/show/movies/sort/" + item + '/' + movie_user_id).then(response => {
                    vm.movies = response.data.Movies;
                    vm.selectedUserID = response.data.SpecificUserID;
                    vm.selectedUser = response.data.SelectedUser;
                });
        },
        likeThisMovie(movie_id, index) {
            if (vm.isLoggedIn) {
                axios.post('/api/rate/movie/'+movie_id+'/1').then(response => {
                   vm.movies.splice(index, 1,  response.data);
                });
            } else {
                alert('Please login in order to like or hate a movie!')
            }

        },
        hateThisMovie(movie_id, index) {
            if (vm.isLoggedIn) {
                axios.post('/api/rate/movie/'+movie_id+'/2').then(response => {
                    vm.movies.splice(index, 1,  response.data);
                });
            } else {
                alert('Please login in order to like or hate a movie!')
            }
        },
        unlikeThisMovie(movie_id, index) {
            if (vm.isLoggedIn) {
                axios.post('/api/unlike/movie/'+movie_id).then(response => {
                    vm.movies.splice(index, 1,  response.data);
                });
            } else {
                alert('Please login in order to like or hate a movie!')
            }
        },
        unhateThisMovie(movie_id, index) {
            if (vm.isLoggedIn) {
                axios.post('/api/unhate/movie/'+movie_id).then(response => {
                    vm.movies.splice(index, 1,  response.data);
                });
            } else {
                alert('Please login in order to like or hate a movie!')
            }
        },
        clearFilter(){
            axios.get("/api/show/movies").then(response => {
                this.movies = response.data;
                this.selectedUser = '';
                this.selectedUserID = 0;
            });
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
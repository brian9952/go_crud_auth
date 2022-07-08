<script>
export default {
    data () {
        return {
            isAuthorized: false,
            data: null,
            post: null,
            errorMessage: null,
        }
    },
    created() {
        this.$watch(
            () => this.$route.params,
            () => {
                this.fetchData()
            },

            { immediate: true }
        )
    },
    methods: {
        fetchData() {
            fetch("http://107.102.183.168:8081/product/showall")
                .then(async response => {
                    const data = await response.json();
                    const isAuthorized = response.headers.get('IsAuthorized');
                    console.log(response);

                    if(!response.ok) {
                        const error = (data && data.message) || response.statusText;
                        return Promise.reject(error);
                    }

                    this.data = data;
                    console.log(this.data);
                    console.log(isAuthorized);
                })
                .catch(error => {
                    this.errorMessage = error;
                    console.error("There was an error !", error);
                });
        }
    },
}
</script>

<template>
    <h1>Products Page</h1>
</template>

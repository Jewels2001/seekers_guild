<script>
const API_URL = `http://localhost:8080`

export default {
    data: () => ({
        entries: null
    }),

    created() {
        // fetch data on init
        this.fetchData()
    },

    methods: {
        async fetchData() {
            const url = `${API_URL}${}`
        }
    }

}
</script>

<template>
    <h1> Leaderboard: </h1>
    <div class="bg-slate-400">
        <table class="border-separate rounded-md border-orange-300 border-spacing-2">
            <tr >
                <th> Name </th>
                <th> Rank </th>
                <th> Current Prestige </th>
                <th><button> Challenge? </button> </th>
            </tr>
            <tr>
                <th> Me </th>
                <th> Office Worker </th>
                <th> 5 </th>
                <th><button> Challenge? </button> </th>
            </tr>
        </table>
    </div>
    
</template>
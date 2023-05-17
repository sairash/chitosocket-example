<script>
    import { goto } from "$app/navigation";
    let message;

    function generateRandomNumber(n) {
        const min = Math.pow(10, n - 1); // Minimum value based on size n
        const max = Math.pow(10, n) - 1; // Maximum value based on size n
        return Math.floor(Math.random() * (max - min + 1)) + min;
    }

    function sendMessage() {
        goto(`/chat/${message}/${generateRandomNumber(3)}`);
        message = "";
    }
</script>

<div
    class="flex justify-center items-center px-10 h-screen antialiased text-gray-800 bg-gray-100"
>
    <form
        on:submit|preventDefault={() => sendMessage()}
        class="flex flex-row items-center h-16 rounded-xl bg-white w-full px-4"
    >
        <div class="flex-grow ml-4">
            <div class="relative w-full">
                <input
                    type="text"
                    placeholder="Room"
                    class="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 h-10"
                    bind:value={message}
                />
            </div>
        </div>
        <div class="ml-4">
            <button
                class="flex items-center justify-center bg-indigo-500 hover:bg-indigo-600 rounded-xl text-white px-4 py-1 flex-shrink-0"
            >
                <span>Join</span>
                <span class="ml-2">
                    <svg
                        class="w-4 h-4 transform rotate-45 -mt-px"
                        fill="none"
                        stroke="currentColor"
                        viewBox="0 0 24 24"
                        xmlns="http://www.w3.org/2000/svg"
                    >
                        <path
                            stroke-linecap="round"
                            stroke-linejoin="round"
                            stroke-width="2"
                            d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
                        />
                    </svg>
                </span>
            </button>
        </div>
    </form>
</div>

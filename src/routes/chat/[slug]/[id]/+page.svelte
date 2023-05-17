<script>
  import { page } from "$app/stores";
  import { onMount } from "svelte";
  import MesLib from "$lib/message.svelte";
  import { chitosocket } from "@sairash/chitosocket";
  const user_id = $page.params.id;
  const user_room = $page.params.slug;

  let cs;
  let message = "";
  let all_messages = [];

  onMount(async () => {
    cs = new chitosocket(`ws://localhost:6969/ws/`);

    cs.on("message", (data) => {
      all_messages = [
        ...all_messages,
        { user_id: data.user_id, msg: data.message, type_of_msg: "message" },
      ];
    });

    cs.on("room_event", (data) => {
      all_messages = [
        ...all_messages,
        { user_id: data.user_id, msg: data.message, type_of_msg: "event" },
      ];
    });

    cs.emit("add_to_room", { room_id: user_room, user_id: user_id });
  });

  function sendMessage() {
    cs.emit("message", { user_id: user_id, message: message });
    message = "";
  }
</script>

<div class="flex h-screen antialiased text-gray-800">
  <div class="flex flex-row h-full w-full overflow-x-hidden">
    <div class="flex flex-col flex-auto h-full">
      <div
        class="flex flex-col flex-auto flex-shrink-0 bg-gray-100 h-full p-4"
      >
        <div class="flex flex-col h-full overflow-x-auto mb-4">
          <div class="flex flex-col h-full">
            <div class="grid grid-cols-12 gap-y-2">
              {#each all_messages as one_message}
                <MesLib {one_message} {user_id} />
              {:else}
                <div
                  class="col-start-1 col-end-12 text-center text-xs text-gray-600"
                >
                  ---- Waiting For Message ----
                </div>
              {/each}
            </div>
          </div>
        </div>
        <form
          on:submit|preventDefault={() => sendMessage()}
          class="flex flex-row items-center h-16 rounded-xl bg-white w-full px-4"
        >
          <div>
            <button
              class="flex items-center justify-center text-gray-400 hover:text-gray-600"
            >
              <svg
                class="w-5 h-5"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
                xmlns="http://www.w3.org/2000/svg"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M15.172 7l-6.586 6.586a2 2 0 102.828 2.828l6.414-6.586a4 4 0 00-5.656-5.656l-6.415 6.585a6 6 0 108.486 8.486L20.5 13"
                />
              </svg>
            </button>
          </div>
          <div class="flex-grow ml-4">
            <div class="relative w-full">
              <input
                type="text"
                class="flex w-full border rounded-xl focus:outline-none focus:border-indigo-300 pl-4 h-10"
                bind:value={message}
              />
              <button
                class="absolute flex items-center justify-center h-full w-12 right-0 top-0 text-gray-400 hover:text-gray-600"
              >
                <svg
                  class="w-6 h-6"
                  fill="none"
                  stroke="currentColor"
                  viewBox="0 0 24 24"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path
                    stroke-linecap="round"
                    stroke-linejoin="round"
                    stroke-width="2"
                    d="M14.828 14.828a4 4 0 01-5.656 0M9 10h.01M15 10h.01M21 12a9 9 0 11-18 0 9 9 0 0118 0z"
                  />
                </svg>
              </button>
            </div>
          </div>
          <div class="ml-4">
            <button
              class="flex items-center justify-center bg-indigo-500 hover:bg-indigo-600 rounded-xl text-white px-4 py-1 flex-shrink-0"
            >
              <span>Send</span>
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
    </div>
  </div>
</div>

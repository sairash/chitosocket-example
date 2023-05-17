<script>
     export let one_message;
     export let user_id;
</script>

{#if one_message.type_of_msg == "message"}
     {#if one_message.user_id == user_id}
          <div class="col-start-6 col-end-13 p-3 rounded-lg">
               <div class="flex items-center justify-start flex-row-reverse">
                    <div
                         class="flex items-center justify-center h-10 w-10 rounded-full bg-indigo-500 flex-shrink-0"
                    >
                         {one_message.user_id}
                    </div>
                    <div
                         class="relative mr-3 text-sm bg-indigo-100 py-2 px-4 shadow rounded-xl"
                    >
                         <div>{one_message.msg}</div>
                    </div>
               </div>
          </div>
     {:else}
          <div class="col-start-1 col-end-8 p-3 rounded-lg">
               <div class="flex flex-row items-center">
                    <div
                         class="flex items-center justify-center h-10 w-10 rounded-full bg-indigo-500 flex-shrink-0"
                    >
                         {one_message.user_id}
                    </div>
                    <div
                         class="relative ml-3 text-sm bg-white py-2 px-4 shadow rounded-xl"
                    >
                         <div>
                              {one_message.msg}
                         </div>
                    </div>
               </div>
          </div>
     {/if}
{:else if one_message.type_of_msg == "event"}
     <div class="col-start-1 col-end-12 text-center text-xs text-gray-600">
          ---- {one_message.msg}
          <span class="font-bold text-rose-500">
               {one_message.user_id}
          </span> ----
     </div>
{/if}

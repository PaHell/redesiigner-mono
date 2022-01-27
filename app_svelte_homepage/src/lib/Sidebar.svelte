<script lang="ts" context="module">
  import Icon from './Icon.svelte';
  import Btn from './Button.svelte';
</script>

<script lang="ts">
  export let title = "Title";
  export let img = "";
  export let visible = false;
  export let colored = false;
  export let right = false;
  
  let display = false;

  $: {
    if (visible) display = true;
    else setTimeout(() => {
      visible = false;
      display = false;
    }, 500);
  }
</script>

<template>
  {#if display}
    <div class="app-sidebar-ct app-sidebar-{right ? 'right' : 'left'}">
      <button class="app-sidebar-bg animate__animated {visible ? 'animate__fadeIn' : 'animate__fadeOut'}"
           on:click={(e) => visible = !visible }>
        <Icon name="cross" css="items-start justify-center animate__animated {visible ? 'animate__bounceIn' : 'animate__bounceOut'}{right ? 'Left' : 'Right'}"/>
      </button>
      <div class="app-sidebar {colored ? 'app-colored' : ''} animate__animated {visible ? 'animate__bounceIn' : 'animate__bounceOut'}{right ? 'Right' : 'Left'}">
        {#if $$slots.header}
          <header>
            <div class="app-top">
              <h2 class="flex-1 text-xl dark:text-white text-gray-900">
                { title }
              </h2>
              <Btn icon="cross"
                   color="{colored ? 'accent' : 'transparent-default'}"
                on:click={(e) => visible = false }/>
            </div>
            <slot name="header"/>
          </header>
        {/if}
        <main>
          <slot>
            <p class="p-4 text-sm dark:text-gray-400 text-gray-500">Default Slot</p>
          </slot>
        </main>
        {#if $$slots.footer}
          <footer>
            <slot name="footer"/>
          </footer>
        {/if}
      </div>
    </div>
  {/if}
</template>

<style global lang="postcss">
    .app-sidebar-ct {
        @apply fixed z-10
        overflow-hidden
        flex justify-center items-stretch
        w-full h-full inset-0;

        & > .app-sidebar-bg {
            @apply fixed w-full h-full
            transition-all
            bg-gray-300 bg-opacity-50
            text-gray-700
            hover:text-gray-900
            active:text-gray-900;

            & > .app-icon {
              @apply w-full h-full p-8
              items-start place-content-end;
            }
        }

        & > .app-sidebar {
            @apply relative flex flex-col w-[32rem]
            bg-gray-100 shadow;

            & > * {
              @apply border-gray-300;
            }

            & > header {
              @apply border-b bg-gray-50;

              & > .app-top {
                @apply flex justify-center items-center p-4;

                & > .app-button {
                  @apply p-2;
                }
              }
            }
            & > main {
              @apply flex-1 overflow-y-auto;
            }
            & > footer {
              @apply border-t;
            }

            &.app-colored > header {
              @apply bg-accent-600 border-accent-700;
              & h2 { @apply text-white; }
              & p { @apply text-accent-300; }
            }
          }

          &.app-sidebar-left {
            @apply justify-start;

            & > .app-sidebar-bg {
              @apply pl-[24rem];

              & > .app-icon {
                @apply justify-start;
              }
            }

            & > .app-sidebar {
              @apply ml-[-8rem];
              & > * { @apply pl-[8rem] border-r; }
            }
        }
          
          &.app-sidebar-right {
            @apply justify-end;

            & > .app-sidebar-bg {
              @apply pr-[24rem];
              & > .app-icon {
                @apply justify-end;
              }
            }
            & > .app-sidebar {
              @apply mr-[-8rem];
              & > * { @apply pr-[8rem] border-l; }
            }
        }
    }

    html.dark .app-sidebar-ct {
        & > .app-sidebar-bg {
            @apply bg-gray-900 bg-opacity-75
            text-gray-500
            hover:text-gray-300
            active:text-gray-300;
        }
        & > .app-sidebar {
            @apply bg-gray-800 shadow-none ring-2 ring-gray-900;

            & > * { @apply border-gray-700; }
            & > header {
              @apply border-b-2 bg-transparent;
            }
            & > footer {
              @apply border-t-2;
            }
            &.app-colored > header {
              @apply bg-accent-700 border-accent-600;
              & p { @apply text-accent-400; }
            }
        }

        &.app-sidebar-left {
          & > .app-sidebar {
            & > * { @apply border-gray-700 border-r-2; }
          }
        }
        &.app-sidebar-right {
          & > .app-sidebar {
            & > * { @apply border-gray-700 border-l-2; }
          }
        }
    }
</style>

<script lang="ts" context="module">
      import { onMount } from 'svelte';
	import Icon from '$lib/Icon.svelte';
</script>

<script lang="ts">
	export let text : string = "";
	export let message : string | undefined;
	export let icon : string | undefined;
	export let color: string | undefined;
	export let bg: string | undefined;
      export let inline: true;
</script>

<template>
      {#if text}
      <div class={`app-alert
            app-alert-color-${color}
            app-alert-bg-${bg}
            ${icon ? 'app-alert-icon' : ''}
            ${inline ? 'app-alert-inline' : ''}`}>
                  {#if icon}
                        <Icon name={icon}/>
                  {/if}
                  <p>
                        <span>{text}</span>
                        {#if message}
                              {#if !inline}
                                    <br/>
                              {/if}
                              <span>{message}</span>
                        {/if}
                  </p>
            </div>
      {/if}
</template>

<style global lang="postcss">
      .app-alert {
            @apply flex items-start
            rounded-md border p-2;

            & > .app-icon {
                  @apply mr-2 text-inherit;
            }

            & > p {
                  @apply text-sm;
                  & > span:first-child {
                        @apply font-medium;
                  }
            }

            &.app-alert-icon {
                  &.app-alert-inline {
                        & > p {
                              @apply leading-ic-base;
                        }
                  }
                  &:not(.app-alert-inline) {
                        & > p > span:first-child {
                              @apply leading-ic-base;
                        }
                  }
            }

            /* colors */

            &.app-alert-color-default {
                  @apply bg-gray-100 border-gray-200;
                  & > .app-icon {
                        @apply text-gray-500;
                  }
                  & > p {
                        @apply text-gray-600;
                        & > span:first-child {
                              @apply text-gray-800;
                        }
                  }
            }
            &.app-alert-color-accent {
                  @apply bg-accent-100 border-accent-200;
                  & > .app-icon {
                        @apply text-accent-500;
                  }
                  & > p {
                        @apply text-accent-600;
                        & > span:first-child {
                              @apply text-accent-800;
                        }
                  }
            }
            &.app-alert-color-success {
                  @apply bg-success-100 border-success-200;
                  & > .app-icon {
                        @apply text-success-500;
                  }
                  & > p {
                        @apply text-success-600;
                        & > span:first-child {
                              @apply text-success-800;
                        }
                  }
            }
            &.app-alert-color-warning {
                  @apply bg-warning-100 border-warning-200;
                  & > .app-icon {
                        @apply text-warning-600;
                  }
                  & > p {
                        @apply text-warning-700;
                        & > span:first-child {
                              @apply text-warning-900;
                        }
                  }
            }
            &.app-alert-color-danger {
                  @apply bg-danger-100 border-danger-200;
                  & > .app-icon {
                        @apply text-danger-500;
                  }
                  & > p {
                        @apply text-danger-600;
                        & > span:first-child {
                              @apply text-danger-800;
                        }
                  }
            }
            &.app-alert-color-info {
                  @apply bg-info-100 border-info-200;
                  & > .app-icon {
                        @apply text-info-500;
                  }
                  & > p {
                        @apply text-info-600;
                        & > span:first-child {
                              @apply text-info-800;
                        }
                  }
            }
      }
</style>

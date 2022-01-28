<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
	import { onMount } from 'svelte';
</script>

<script lang="ts">
	export let text: string | undefined = undefined;
	export let icon: string | undefined = undefined;
	export let color: string | undefined = undefined;
	export let bg: string | undefined = undefined;
	export let disabled: boolean = false;
	export let isLoading: boolean = true;

	let ref: HTMLButtonElement | undefined;

	onMount(() => {
		ref?.focus();
	});
</script>

<template>
	<button on:click bind:this={ref}
		style="outline: none;"
		class="app-button app-button-bg-{bg ?? 'default'} app-button-color-{color ?? 'default'}">
		{#if isLoading}
			<div class="app-loading fixed"></div>
		{/if}
		{#if icon}
			<Icon name={icon} />
		{/if}
		{#if text}
			<p>{text}</p>
		{/if}
	</button>
</template>

<style global lang="postcss">
	.app-bar-button {
		@apply flex
        flex-col-reverse sm:flex-row
        sm:justify-end
        sm:p-3 p-4
        bg-gray-50;

		&.app-bar-reverse {
			@apply sm:justify-start;
		}

		& .app-button {
			&:not(:last-child) {
				@apply mt-2 sm:mt-0 mr-0 sm:mr-2;
			}
		}

		& > p {
			@apply text-gray-500 px-2 self-center;
		}
	}

	.app-button {
		@apply flex items-center
        justify-center sm:justify-start
        px-4 py-2 border rounded-md
        text-base sm:text-sm font-medium
        transition-all;

		& > p {
            @apply mb-[1px] font-medium;
            &:not(:first-child) {
                @apply ml-1;
            }
		}

		& > *:not(.app-button-default) {
			@apply border-transparent;
		}

        /* ring for non-transparent */

        &:not(.app-button-bg-transparent) {
            @apply focus:ring-2 focus:ring-offset-2
            focus:ring-offset-gray-50 focus:ring-accent-600;

            &.app-button-bg-success,
            &.app-button-bg-default.app-button-color-success {
                @apply focus:ring-success-600;
            }

            &.app-button-bg-warning,
            &.app-button-bg-default.app-button-color-warning {
                @apply focus:ring-warning-600;
            }

            &.app-button-bg-danger,
            &.app-button-bg-default.app-button-color-danger {
                @apply focus:ring-danger-600;
            }

            &.app-button-bg-info,
            &.app-button-bg-default.app-button-color-info {
                @apply focus:ring-info-600;
            }
        }

        /* colored text */
		&.app-button-bg-default,
        &.app-button-bg-transparent {
            &.app-button-color-default {
                & > p {
                    @apply text-gray-700;
                }
                & > .app-icon {
                    @apply text-gray-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-gray-900;
                    }
                    & > .app-icon {
                        @apply text-gray-700;
                    }
                }
            }

            &.app-button-color-accent {
                & > p {
                    @apply text-accent-600;
                }
                & > .app-icon {
                    @apply text-accent-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-accent-700;
                    }
                    & > .app-icon {
                        @apply text-accent-600;
                    }
                }
            }

            &.app-button-color-success {
                & > p {
                    @apply text-success-600;
                }
                & > .app-icon {
                    @apply text-success-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-success-700;
                    }
                    & > .app-icon {
                        @apply text-success-600;
                    }
                }
            }
            
            &.app-button-color-warning {
                & > p {
                    @apply text-warning-600;
                }
                & > .app-icon {
                    @apply text-warning-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-warning-700;
                    }
                    & > .app-icon {
                        @apply text-warning-600;
                    }
                }
            }
            
            
            &.app-button-color-danger {
                & > p {
                    @apply text-danger-600;
                }
                & > .app-icon {
                    @apply text-danger-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-danger-700;
                    }
                    & > .app-icon {
                        @apply text-danger-600;
                    }
                }
            }

            &.app-button-color-info {
                & > p {
                    @apply text-info-600;
                }
                & > .app-icon {
                    @apply text-info-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-info-700;
                    }
                    & > .app-icon {
                        @apply text-info-600;
                    }
                }
            }
        }

        &.app-button-bg-default {
			@apply bg-gray-50 hover:bg-white active:bg-gray-100 
            shadow-sm;

            &.app-button-color-default {
                & > p {
                    @apply text-gray-700;
                }
                & > .app-icon {
                    @apply text-gray-500;
                }
                &:hover,
                &:focus,
                &:active {
                    & > p {
                        @apply text-gray-900;
                    }
                    & > .app-icon {
                        @apply text-gray-700;
                    }
                }
            }
        }

        &.app-button-bg-transparent {
			@apply bg-transparent border-transparent
            hover:bg-gray-100 hover:border-gray-300
            active:bg-gray-200 active:border-gray-300;
		}

        /* white text on colored buttons only */
        &:not(.app-button-bg-default):not(.app-button-bg-transparent) {
            & > p,
            & > .app-icon {
                @apply text-white;
            }
            &:hover,
            &:focus,
            &:active {
                & > p,
                & > .app-icon {
                    @apply text-white;
                }
            }
        }

        &.app-button-bg-accent {
			@apply bg-accent-500 border-accent-500
            hover:bg-accent-600 hover:border-accent-700
            active:bg-accent-700 active:border-accent-700;
		}

        &.app-button-bg-success {
			@apply bg-success-500 border-success-500
            hover:bg-success-600 hover:border-success-700
            active:bg-success-700 active:border-success-700;
		}

        &.app-button-bg-warning {
			@apply bg-warning-500 border-warning-500
            hover:bg-warning-600 hover:border-warning-700
            active:bg-warning-700 active:border-warning-700;
		}

        &.app-button-bg-danger {
			@apply bg-danger-500 border-danger-500
            hover:bg-danger-600 hover:border-danger-700
            active:bg-danger-700 active:border-danger-700;
		}

        &.app-button-bg-info {
			@apply bg-info-500 border-info-500
            hover:bg-info-600 hover:border-info-700
            active:bg-info-700 active:border-info-700;
		}
	}
    
	html.dark {
        & .app-bar-button {
            @apply bg-transparent;
            
			& > p {
                @apply text-gray-400;
			}
		}
        & .app-button {
            @apply border-2;

            /* ring for non-transparent */

            &:not(.app-button-bg-transparent) {
                @apply focus:ring-offset-gray-800;

                &.app-button-bg-success,
                &.app-button-bg-default.app-button-color-success {
                    @apply focus:ring-success-500;
                }

                &.app-button-bg-warning,
                &.app-button-bg-default.app-button-color-warning {
                    @apply focus:ring-warning-500;
                }

                &.app-button-bg-danger,
                &.app-button-bg-default.app-button-color-danger {
                    @apply focus:ring-danger-500;
                }

                &.app-button-bg-info,
                &.app-button-bg-default.app-button-color-info {
                    @apply focus:ring-info-500;
                }
            }

            /* colored text */
            &.app-button-bg-default,
            &.app-button-bg-transparent {
                &.app-button-color-default {
                    & > p {
                        @apply text-gray-200;
                    }
                    & > .app-icon {
                        @apply text-gray-300;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p {
                            @apply text-gray-100;
                        }
                        & > .app-icon {
                            @apply text-gray-200;
                        }
                    }
                }

                &.app-button-color-accent {
                    & > p,
                    & > .app-icon {
                        @apply text-accent-300;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p,
                        & > .app-icon {
                            @apply text-accent-300;
                        }
                    }
                }

                &.app-button-color-success {
                    & > p,
                    & > .app-icon {
                        @apply text-success-300;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p,
                        & > .app-icon {
                            @apply text-success-300;
                        }
                    }
                }
                
                &.app-button-color-warning {
                    & > p,
                    & > .app-icon {
                        @apply text-warning-300;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p,
                        & > .app-icon {
                            @apply text-warning-300;
                        }
                    }
                }
                
                
                &.app-button-color-danger {
                    & > p,
                    & > .app-icon {
                        @apply text-danger-300;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p,
                        & > .app-icon {
                            @apply text-danger-300;
                        }
                    }
                }

                &.app-button-color-info {
                    & > p,
                    & > .app-icon {
                        @apply text-info-300;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p,
                        & > .app-icon {
                            @apply text-info-300;
                        }
                    }
                }
            }

            &.app-button-bg-default {
                @apply bg-gray-700 border-gray-700
                hover:bg-gray-700 hover:border-gray-600
                active:bg-gray-600 active:border-gray-600
                shadow-none;

                &.app-button-color-default {
                    & > p {
                        @apply text-gray-100;
                    }
                    & > .app-icon {
                        @apply text-gray-200;
                    }
                    &:hover,
                    &:focus,
                    &:active {
                        & > p {
                            @apply text-white;
                        }
                        & > .app-icon {
                            @apply text-gray-100;
                        }
                    }
                }
            }

            &.app-button-bg-transparent {
                @apply
                hover:bg-gray-800 hover:border-gray-700
                active:bg-gray-700 active:border-gray-700;
            }

            &.app-button-bg-accent {
                @apply bg-accent-600 border-accent-600
                hover:bg-accent-600 hover:border-accent-500
                active:bg-accent-500 active:border-accent-500;
            }

            &.app-button-bg-success {
                @apply bg-success-600 border-success-600
                hover:bg-success-600 hover:border-success-500
                active:bg-success-500 active:border-success-500;
            }

            &.app-button-bg-warning {
                @apply bg-warning-600 border-warning-600
                hover:bg-warning-600 hover:border-warning-500
                active:bg-warning-500 active:border-warning-500;
            }

            &.app-button-bg-danger {
                @apply bg-danger-600 border-danger-600
                hover:bg-danger-600 hover:border-danger-500
                active:bg-danger-500 active:border-danger-500;
            }

            &.app-button-bg-info {
                @apply bg-info-600 border-info-600
                hover:bg-info-600 hover:border-info-500
                active:bg-info-500 active:border-info-500;
            }
        }
	}
</style>

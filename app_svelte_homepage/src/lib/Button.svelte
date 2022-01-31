<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
	import LoadingSpinner from './LoadingSpinner.svelte';
	import { createEventDispatcher, onMount } from 'svelte';
</script>

<script lang="ts">
	export let text: string | undefined = undefined;
	export let icon: string | undefined = undefined;
	export let color: string | undefined = undefined;
	export let bg: string | undefined = undefined;
	export let disabled: boolean = false;
	export let loading: boolean = false;
	export let css: string | undefined = undefined;

	let ref: HTMLButtonElement | undefined;
	const dispatch = createEventDispatcher();

	onMount(() => {
		ref?.focus();
	});

	function onClick(e: Event): void {
		dispatch('click');
		disabled = true;
		loading = true;
		setTimeout(() => {
			disabled = false;
			loading = false;
		}, 750);
	}
</script>

<template>
	<button
		on:click={onClick}
		bind:this={ref}
		{disabled}
		style="outline: none;"
		class={`
            ${css}
            app-button
            app-button-bg-${bg ?? 'default'}
            app-button-color-${color ?? 'default'}
        `}
	>
		{#if loading}
			<div class="animate__animated animate__bounceIn">
				<LoadingSpinner />
			</div>
		{/if}
		<main>
			{#if icon}
				<Icon name={icon} />
			{/if}
			<slot/>
			{#if text}
				<p>{text}</p>
			{/if}
		</main>
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
		@apply flex flex-col items-stretch
        px-4 py-2 border rounded-md
        text-base sm:text-sm font-medium
        transition-all select-none;

		&:disabled {
			@apply opacity-40 pointer-events-none;
			& > main {
				@apply opacity-40;
			}
		}

		&.app-button-no-ring {
			@apply ring-0 ring-offset-0 !important;
		}

		& > div {
			@apply flex justify-center items-center
            h-0 relative top-[.65rem];
		}

		& > main {
			@apply flex transition-opacity;
			& > p {
				@apply mt-[-1.25px] font-medium;
				&:not(:first-child) {
					@apply ml-1;
				}
				/* mirror icon spacing to right if icon exists */
				&:last-child:not(:first-child) {
					@apply pr-1;
				}
			}
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

		/* white ring, colored ring inset*/
		&.app-button-bg-white {
			&:disabled {
				@apply opacity-100;
				& > main {
					@apply opacity-20;
				}
			}
			&.app-button-color-default {
				& .app-loading-spinner {
					@apply stroke-accent-500 !important;
				}
			}
			&.app-button-color-accent {
				@apply border-accent-600 focus:ring-accent-400 focus:ring-offset-accent-500;
			}
			&.app-button-color-success {
				@apply border-success-600 focus:ring-success-400 focus:ring-offset-success-500;
			}
			&.app-button-color-warning {
				@apply border-warning-600 focus:ring-warning-400 focus:ring-offset-warning-500;
			}
			&.app-button-color-danger {
				@apply border-danger-600 focus:ring-danger-400 focus:ring-offset-danger-500;
			}
			&.app-button-color-info {
				@apply border-info-600 focus:ring-info-400 focus:ring-offset-info-500;
			}
		}

		/* colored text */
		&.app-button-bg-white,
		&.app-button-bg-default,
		&.app-button-bg-transparent {
			&.app-button-color-accent {
				& .app-loading-spinner {
					@apply stroke-accent-500;
				}
				& p {
					@apply text-accent-600;
				}
				& .app-icon {
					@apply text-accent-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-accent-700;
					}
					& .app-icon {
						@apply text-accent-600;
					}
				}
			}

			&.app-button-color-success {
				& .app-loading-spinner {
					@apply stroke-success-500;
				}
				& p {
					@apply text-success-600;
				}
				& .app-icon {
					@apply text-success-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-success-700;
					}
					& .app-icon {
						@apply text-success-600;
					}
				}
			}

			&.app-button-color-warning {
				& .app-loading-spinner {
					@apply stroke-warning-500;
				}
				& p {
					@apply text-warning-600;
				}
				& .app-icon {
					@apply text-warning-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-warning-700;
					}
					& .app-icon {
						@apply text-warning-600;
					}
				}
			}

			&.app-button-color-danger {
				& .app-loading-spinner {
					@apply stroke-danger-500;
				}
				& p {
					@apply text-danger-600;
				}
				& .app-icon {
					@apply text-danger-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-danger-700;
					}
					& .app-icon {
						@apply text-danger-600;
					}
				}
			}

			&.app-button-color-info {
				& .app-loading-spinner {
					@apply stroke-info-500;
				}
				& p {
					@apply text-info-600;
				}
				& .app-icon {
					@apply text-info-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-info-700;
					}
					& .app-icon {
						@apply text-info-600;
					}
				}
			}
		}

		&.app-button-bg-white,
		&.app-button-bg-default {
			@apply bg-gray-50 hover:bg-white active:bg-gray-100 
            shadow-sm hover:shadow;

			&.app-button-color-default {
				& p {
					@apply text-gray-700;
				}
				& .app-icon {
					@apply text-gray-500;
				}
				&:hover,
				&:focus,
				&:active {
					& p {
						@apply text-gray-900;
					}
					& .app-icon {
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
		&.app-button-bg-accent,
		&.app-button-bg-success,
		&.app-button-bg-warning,
		&.app-button-bg-danger,
		&.app-button-bg-info {
			& p,
			& .app-icon {
				@apply text-white;
			}
			&:hover,
			&:focus,
			&:active {
				& p,
				& .app-icon {
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

			& p {
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

			/* white ring, colored inset*/
			&.app-button-bg-white {
				@apply border;
				&.app-button-color-default {
					@apply border-gray-900;
				}
				&.app-button-color-accent {
					@apply border-accent-700 focus:ring-accent-500 focus:ring-offset-accent-600;
				}
				&.app-button-color-success {
					@apply border-success-700 focus:ring-success-500 focus:ring-offset-success-600;
				}
				&.app-button-color-warning {
					@apply border-warning-700 focus:ring-warning-500 focus:ring-offset-warning-600;
				}
				&.app-button-color-danger {
					@apply border-danger-700 focus:ring-danger-500 focus:ring-offset-danger-600;
				}
				&.app-button-color-info {
					@apply border-info-700 focus:ring-info-500 focus:ring-offset-info-600;
				}
			}

			/* colored text */
			&.app-button-bg-default,
			&.app-button-bg-transparent {
				&.app-button-color-default {
					& p {
						@apply text-gray-200;
					}
					& .app-icon {
						@apply text-gray-300;
					}
					&:hover,
					&:focus,
					&:active {
						& p {
							@apply text-gray-100;
						}
						& .app-icon {
							@apply text-gray-200;
						}
					}
				}

				&.app-button-color-accent {
					& p,
					& .app-icon {
						@apply text-accent-300;
					}
					&:hover,
					&:focus,
					&:active {
						& p,
						& .app-icon {
							@apply text-accent-300;
						}
					}
				}

				&.app-button-color-success {
					& p,
					& .app-icon {
						@apply text-success-300;
					}
					&:hover,
					&:focus,
					&:active {
						& p,
						& .app-icon {
							@apply text-success-300;
						}
					}
				}

				&.app-button-color-warning {
					& p,
					& .app-icon {
						@apply text-warning-300;
					}
					&:hover,
					&:focus,
					&:active {
						& p,
						& .app-icon {
							@apply text-warning-300;
						}
					}
				}

				&.app-button-color-danger {
					& p,
					& .app-icon {
						@apply text-danger-300;
					}
					&:hover,
					&:focus,
					&:active {
						& p,
						& .app-icon {
							@apply text-danger-300;
						}
					}
				}

				&.app-button-color-info {
					& p,
					& .app-icon {
						@apply text-info-300;
					}
					&:hover,
					&:focus,
					&:active {
						& p,
						& .app-icon {
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
					& p {
						@apply text-gray-100;
					}
					& .app-icon {
						@apply text-gray-200;
					}
					&:hover,
					&:focus,
					&:active {
						& p {
							@apply text-white;
						}
						& .app-icon {
							@apply text-gray-100;
						}
					}
				}
			}

			&.app-button-bg-transparent {
				@apply hover:bg-gray-800 hover:border-gray-700
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

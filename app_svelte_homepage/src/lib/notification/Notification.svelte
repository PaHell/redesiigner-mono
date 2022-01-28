<script lang="ts" context="module">
	import Icon from '$lib/Icon.svelte';
	import Btn from '$lib/Button.svelte';
	import { modalStore } from '$lib/../store';
	import { createEventDispatcher, onMount, SvelteComponent } from 'svelte';

	const variantClose = {
		undefined: {
			icon: 'cross',
			text: null,
			color: 'transparent-default'
		},
		default: {
			icon: 'cross',
			text: null,
			color: 'transparent-default'
		},
		button: {
			icon: 'cross',
			text: 'Dismiss',
			color: 'default'
		},
		reply: {
			icon: null,
			text: 'Dismiss',
			color: 'transparent-default'
		}
	};
</script>

<script lang="ts">
	export let icon: string | undefined = undefined;
	export let img: string | undefined = undefined;
	export let title = 'Title';
	export let text = 'Text';
	export let color = 'accent';
	export let variant = 'default';
	export let buttons: TButton<void, void>[] = [];
	export let remove = () => {};

	let visible = true;
	function close(): void {
		visible = false;
		setTimeout(() => {
			visible = true;
			remove();
		}, 500);
	}
</script>

<template>
	<div class="app-nt-wrapper" class:app-visible={visible}>
		<menu
			class="app-nt app-nt-variant-{variant} app-nt-color-{color} animate__animated {visible
				? 'animate__jackInTheBox'
				: 'animate__rollOut'}"
		>
			<header>
				{#each buttons as button}
					<Btn
						text={button.text}
						icon={button.icon}
						color={button.color}
						on:click={() => button.OnClick(close)}
					/>
				{/each}
				<Btn
					icon={variantClose[variant]?.icon}
					color={variantClose[variant]?.color}
					text={variantClose[variant]?.text}
					on:click={() => close()}
				/>
			</header>
			<main>
				{#if img}
					<div class="app-img">
						<img src={img} alt={title} />
					</div>
				{/if}
				{#if icon}
					<Icon name={icon} css="text-{color}-500" large />
				{/if}
				<div>
					<h6 class="text-lg sm:text-base font-semibold text-gray-800 dark:text-white">
						{title}
					</h6>
					<p class="text-base sm:text-sm text-gray-500 dark:text-gray-400">
						{text}
					</p>
					<slot />
				</div>
			</main>
		</menu>
	</div>
</template>

<style global lang="postcss">
	.app-nt-wrapper {
		@apply max-h-0 overflow-visible;
		transition: max-height 0.4s;
		animation-duration: 0.4s;

		&.app-visible {
			@apply max-h-48;
		}
		&:not(:last-child) > .app-nt {
			@apply sm:mb-2;
		}
	}

	.app-nt {
		@apply overflow-hidden w-full
        p-4 py-6 sm:py-4
        rounded-none sm:rounded-lg transition-all
        bg-gray-50 border border-gray-300 shadow;

		& > header {
			@apply flex justify-end items-center;
		}

		& > main {
			@apply flex items-start;

			& > .app-icon {
				@apply mr-4;

				& + div > h6 {
					@apply leading-icon sm:leading-icon;
				}
			}

			& > .app-img {
				@apply flex flex-wrap justify-center items-center
                w-11 h-11 mr-4
                rounded-full overflow-hidden;

				&:after {
					content: '';
					@apply block absolute
                    w-11 h-11 rounded-full;
					box-shadow: inset 0 0 4px rgba(0, 0, 0, 0.3), inset 0 0 12px 1px rgba(0, 0, 0, 0.2);
				}

				& > img {
					@apply min-w-full min-h-full;
				}
			}
		}

		/* Colors */

		&.app-nt-color-default,
		&.app-nt-color-undefined {
			& > main > .app-icon {
				@apply text-gray-500;
			}
		}

		/* Variants */

		&.app-nt-variant-default,
		&.app-nt-variant-undefined {
			& > header {
				@apply h-8 -mb-8 overflow-visible
                relative sm:-top-2 -top-3 left-2;

				& > .app-button {
					@apply p-2 focus:ring-0;
					box-shadow: none !important;
					& > p:first-child:last-child {
						@apply sm:px-0 px-2;
					}
				}
			}
		}

		&.app-nt-variant-button {
			@apply flex flex-col-reverse;

			& > header {
				@apply mt-3 ml-[32px] sm:ml-icon pl-4 justify-start;

				& > .app-button:not(:last-child) {
					@apply mr-2;
				}
			}
		}

		&.app-nt-variant-reply {
			@apply flex flex-row-reverse p-0 overflow-visible;

			& > header {
				@apply flex-col items-stretch m-[-1px];

				& > .app-button {
					@apply flex-1 justify-center rounded-none
                    border-gray-300;
					box-shadow: none !important;

					&:hover,
					&:hover,
					&:hover {
						@apply relative;
					}

					&:first-child {
						@apply sm:rounded-tr-lg;
					}

					&:last-child {
						@apply sm:rounded-br-lg;
					}

					&:not(:last-child) {
						@apply mb-[-.5px] mr-0;
					}
				}
			}

			& > main {
				@apply flex-1 p-4;
			}
		}
	}

	html.dark .app-nt-wrapper {
		& > .app-nt {
			@apply shadow-none
            bg-gray-800 border-gray-700
            sm:border-2 border-b-2
            sm:ring-2 sm:ring-gray-900;

			& > header {
				& > .app-button > p {
					@apply font-bold;
				}
			}

			& > main {
			}

			/* Colors */

			&.app-nt-color-default,
			&.app-nt-color-undefined {
				& > main > .app-icon {
					@apply text-gray-400;
				}
			}

			/* Variants */

			&.app-nt-variant-reply {
				& > header {
					@apply m-[-2px];
					& > .app-button {
						@apply py-0
                        border-gray-700 hover:border-gray-600 active:border-gray-600;

						&:not(:last-child) {
							@apply mb-[-2px];
						}
					}
				}
			}
		}

		&:last-child > .app-nt {
			@apply border-b-2 shadow-[0_2px_0_0_rgba(0,0,0,0.3)];
		}
	}
</style>

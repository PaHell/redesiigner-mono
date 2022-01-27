import { writable } from 'svelte/store';

const SOUND_NOTIFICATION = "window_3";
const SOUND_BANNER = "window_2";
const SOUND_MODAL = "window_1";

export const soundStore = function () {
	const { subscribe, set } = writable<string>("");
	return {
		subscribe,
		play: (sound: string) => {
			set("");
			set(sound);
		},
	};
}();

export const bannerStore = function () {
	const { subscribe, set, update } = writable<(TBanner|null)[]>([]);
	return {
		subscribe,
		add: (banner: TBanner) => {
			soundStore.play(SOUND_BANNER);
			update(value => {
				const last = value.push(null) - 1;
				// deep copy
				value[last] = {
					...banner,
					buttons: [...banner.buttons],
					remove: () => update(all => {
						all[last] = null;
						return all;
					}),
				}
                return value;
            });
		},
        clear: () => {
			set([]);
		}
	};
}();

export const notificationStore = function () {
	const { subscribe, set, update } = writable<(TNotification|null)[]>([]);
	return {
		subscribe,
		add: (notification: TNotification) => {
			soundStore.play(SOUND_NOTIFICATION);
			update(value => {
				const last = value.push(null) - 1;
				// deep copy
				value[last] = {
					...notification,
					buttons: [...notification.buttons],
					remove: () => update(all => {
						all[last] = null;
						return all;
					}),
				}
                return value;
            });
		},
        clear: () => {
			set([]);
		}
	};
}();

export const modalStore = function () {
	const { subscribe, set, update } = writable<(TModal|null)[]>([]);
	return {
		subscribe,
		add: (modal: TModal) => {
			update(value => {
				const last = value.push(null) - 1;
				// deep copy
				value[last] = {
					...modal,
					buttons: [...modal.buttons],
					remove: () => update(all => {
						all[last] = null;
						return all;
					}),
				}
                return value;
            });
		},
        clear: () => {
			set([]);
		}
	};
}();
/// <reference types="@sveltejs/kit" />

type TValidation = {
	regex: RegExp;
	message: string;
}

type TButton<TIn, TOut> = {
	icon?: string | undefined;
	text?: string | undefined;
	color?: string | undefined;
	bg?: string | undefined;
	OnClick: (TIn) => TOut;
};

type TBanner = {
	icon: string;
	text: string;
	color?: string;
	buttons: Button<void, void>[];
	remove?: () => void;
};

type TModal = {
	icon: string;
	title: string;
	text: string;
	color: string;
	buttons: Button<(close: () => void) => void, void>[];
	remove?: () => void;
};

type TNotification = {
	icon?: string;
	img?: string;
	title: string;
	text: string;
	color?: string;
	variant?: string;
	buttons: Button<void, void>[];
	remove?: () => void;
};

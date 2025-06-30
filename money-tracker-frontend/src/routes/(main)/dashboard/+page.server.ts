import { fail, redirect, type Actions } from '@sveltejs/kit';
import type { PageServerLoad } from './$types';

export const load = (async ({ cookies }) => {
	return {};
}) satisfies PageServerLoad;

export const actions: Actions = {
	verify: async ({ fetch, request }) => {
		const form = await request.formData();
		const token = form.get('otp') as string;
		// const username = form.get('username') as string;

		if (token === '') {
			return fail(400, { message: 'Dont Forget Your OTP' });
		}
		const body = JSON.stringify({ token });

		console.log(body);
		try {
			const response = await fetch(`http://127.0.0.1:8080/api/v1/user/verify?token=${token}`, {
				method: 'PATCH',
				headers: {
					'Content-Type': 'application/json',
					Accept: 'application/json'
				}
			});

			const result = await response.json();
			console.log(result);

			if (!response.ok) {
				if (result.details === 'redis value didnt exists') {
					return fail(400, { message: 'Are you sure that is the correct OTP?' });
				}
				return fail(response.status, result);
			}

			return {
				result
			};
		} catch (error) {
			return fail(500, { message: 'Server error occurred' });
		}
	}
};

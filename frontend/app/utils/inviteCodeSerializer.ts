import { type InviteCode, type FormattedInviteCode } from "~/types/inviteCode"

export default (inviteCode: InviteCode): FormattedInviteCode =>
(
	{
		id: inviteCode.id,
		code: inviteCode.id,
		uses: inviteCode.uses,
		maxUses: inviteCode.max_uses,
		createdAt: inviteCode.created_at
	}
);
package conversationclientv2

import "fmt"

// This code was automatically generated by github.com/mittwald/api-client-go-builder.
// DO NOT EDIT.

// This data type was generated from the following JSON schema:
// type: "string"
// enum:
//    - "createdAt"
//    - "lastMessage.createdAt"
//    - "title"
//    - "priority"
//    - "shortId"
//    - "conversationId"

type ListConversationsRequestQuerySortItem string

const ListConversationsRequestQuerySortItemCreatedAt ListConversationsRequestQuerySortItem = "createdAt"
const ListConversationsRequestQuerySortItemLastMessagecreatedAt ListConversationsRequestQuerySortItem = "lastMessage.createdAt"
const ListConversationsRequestQuerySortItemTitle ListConversationsRequestQuerySortItem = "title"
const ListConversationsRequestQuerySortItemPriority ListConversationsRequestQuerySortItem = "priority"
const ListConversationsRequestQuerySortItemShortID ListConversationsRequestQuerySortItem = "shortId"
const ListConversationsRequestQuerySortItemConversationID ListConversationsRequestQuerySortItem = "conversationId"

func (e ListConversationsRequestQuerySortItem) Validate() error {
	if e == ListConversationsRequestQuerySortItemCreatedAt || e == ListConversationsRequestQuerySortItemLastMessagecreatedAt || e == ListConversationsRequestQuerySortItemTitle || e == ListConversationsRequestQuerySortItemPriority || e == ListConversationsRequestQuerySortItemShortID || e == ListConversationsRequestQuerySortItemConversationID {
		return nil
	}
	return fmt.Errorf("unexpected value for type %T: %s", e, e)
}

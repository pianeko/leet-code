/**
 * Definition for singly-linked list.
 * class ListNode {
 *     val: number
 *     next: ListNode | null
 *     constructor(val?: number, next?: ListNode | null) {
 *         this.val = (val===undefined ? 0 : val)
 *         this.next = (next===undefined ? null : next)
 *     }
 * }
 */

class ListNode {
  val: number
  next: ListNode | null
  constructor(val?: number, next?: ListNode | null) {
    this.val = (val===undefined ? 0 : val)
    this.next = (next===undefined ? null : next)
  }
}

const l1 = {
  val: 2,
  next: {
    val: 4,
    next: {
      val: 3,
      next: null
    }
  } 
}
// l1 = [2,4,3]
// l2 = [5,6,4]
// output = [7,0,8]

function addTwoNumbers(l1: ListNode | null, l2: ListNode | null): ListNode | null {
  const l1Array: number[] = [];
  const l2Array: number[] = [];
  const outputArray: number[] = [];
  let carry = 0;
  let outputListNode: ListNode | null = null;

  // initialize
  l1Array.push(l1.val);
  l2Array.push(l2.val);

  // convert to array
  while(l1.next !== null || l2.next !== null) {
    l1 = (l1.next===null ? new ListNode() : l1.next);
    l2 = (l2.next===null ? new ListNode() : l2.next);

    l1Array.push(l1.val);
    l2Array.push(l2.val);
  }

  // caluculate
  l1Array.forEach((val: number, index: number) => {
    let valTmp = val + l2Array[index] + carry;
    carry = 0;
    if (valTmp >= 10) {
      valTmp = valTmp - 10;
      carry = 1;
    }
    outputArray.push(valTmp);
  })
  if (carry === 1) {
    outputArray.push(1);
  }

  // convert to ListNode
  for (let i = outputArray.length - 1; i >= 0; i--) {
    outputListNode = new ListNode(outputArray[i], outputListNode); 
  }

  return outputListNode;
};
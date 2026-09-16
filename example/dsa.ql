;; ql code for matching all dsa notes

(unless (tag "not-dsa")
    (or
        ;;
        ;; platform tags
        ;;
        (tag "leetcode")
        (tag "codeforces")
        ;;
        ;; difficulty tags
        ;;
        (tag "easy")
        (tag "medium")
        (tag "hard")))

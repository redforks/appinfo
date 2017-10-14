package appinfo_test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/redforks/appinfo"
	"github.com/redforks/osutil"
	"github.com/redforks/testing/iotest"
	"github.com/redforks/testing/reset"
)

var _ = Describe("appinfo", func() {
	var (
		tempDir iotest.TempTestDir
	)

	BeforeEach(func() {
		reset.Enable()

		tempDir = iotest.NewTempTestDir()
		Ω(os.Setenv("_root_dir", tempDir.Dir())).Should(Succeed())
	})

	AfterEach(func() {
		_ = os.Setenv("_root_dir", "")
		reset.Disable()
	})

	It("SetInfo", func() {
		appinfo.SetInfo("foo", "abcdef")
		Ω(appinfo.CodeName()).Should(Equal("foo"))
		Ω(appinfo.Version()).Should(Equal("abcdef"))
	})

	It("Generate InstallID", func() {
		appinfo.SetInfo("foo", "ver")
		Ω(appinfo.InstallID()).ShouldNot(Equal(""))
	})

	It("Restore InstallID", func() {
		idFile := filepath.Join(tempDir.Dir(), "var/lib/spork/foo.id")
		Ω(osutil.WriteFile(idFile, []byte("3456"), 0700, 0700)).Should(Succeed())
		appinfo.SetInfo("foo", "ver")
		Ω(appinfo.InstallID()).Should(Equal("3456"))
	})

})

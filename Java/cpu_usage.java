
import com.sun.management.OperatingSystemMXBean;
import java.lang.management.ManagementFactory;
import java.util.concurrent.TimeUnit;
import java.text.DecimalFormat;

public class cpu_usage {

    public static void main(String[] args) throws InterruptedException {
        OperatingSystemMXBean operatingSystemMXBean =(com.sun.management.OperatingSystemMXBean) ManagementFactory.getOperatingSystemMXBean();

        System.out.println("System Cpu Usage:");
        while(true) {
            DecimalFormat f = new DecimalFormat("0.00");
            Double d = operatingSystemMXBean.getSystemCpuLoad();
            System.out.println(f.format(d)+"%");
            TimeUnit.SECONDS.sleep(2);
        }
    }
}
